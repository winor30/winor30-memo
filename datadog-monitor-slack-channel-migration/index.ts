import { client, v1 as datadog } from '@datadog/datadog-api-client';

if (!process.env.DATADOG_API_KEY || !process.env.DATADOG_APP_KEY) {
  throw new Error('DATADOG_API_KEY and DATADOG_APP_KEY must be set');
}

const configuration = client.createConfiguration({
  authMethods: {
    apiKeyAuth: process.env.DATADOG_API_KEY || '',
    appKeyAuth: process.env.DATADOG_APP_KEY || '',
  },
});

const apiInstance = new datadog.MonitorsApi(configuration);

const REPLACE_SLACK_CHANNEL_MAP = {
  "channel_1_prev": "channel_1_new",
  "channel_2_prev": "channel_2_new",
};

// モニターの一覧を取得する関数
const fetchMonitors = async (): Promise<datadog.Monitor[]> => {
  try {
    const monitors = await apiInstance.listMonitors({
      pageSize: 100,
    });
    return monitors || [];
  } catch (error) {
    console.error('Error fetching monitors:', error);
    throw error;
  }
};

// Slack通知チャンネルを更新する関数
const updateSlackChannel = async (monitorId: number) => {
  try {
    const monitor = await apiInstance.getMonitor({ monitorId });

    const message = monitor.message;
    if (!message) {
      return;
    }

    const hasTargetSlackChannel = Object.keys(REPLACE_SLACK_CHANNEL_MAP).some((previousSlackChannel) => message.includes(previousSlackChannel));

    if (!hasTargetSlackChannel) {
      return;
    }

    const newMessage = Object.keys(REPLACE_SLACK_CHANNEL_MAP).reduce((acc, previousSlackChannel) => {
      const newSlackChannel = REPLACE_SLACK_CHANNEL_MAP[previousSlackChannel as keyof typeof REPLACE_SLACK_CHANNEL_MAP];
      return acc.replace(previousSlackChannel, newSlackChannel);
    }, message);

    const updateParams = {
      monitorId: monitorId,
      body: {
        message: newMessage,
      },
    };

    await apiInstance.updateMonitor(updateParams);

    console.log(`Updated monitor ${monitorId} with new Slack channel.`);
    console.log(message);
    console.log(newMessage);

  } catch (error) {
    console.error(`Error updating monitor ${monitorId}:`, error);
  }
};

// メイン処理
const main = async () => {
  try {
    const monitors = await fetchMonitors();
    for (const monitor of monitors) {
      if (monitor.id) {
        await updateSlackChannel(monitor.id);
      }
    }
    console.log('Slack channels updated successfully.');
  } catch (error) {
    console.error('Error updating monitors:', error);
  }
};

main();
