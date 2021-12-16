package structure.adapter;

public class SlackNotifierAdapter implements Notifier {
    private final SlackClient slackClient;
    private final String receivingChannel;

    SlackNotifierAdapter(String slackApiKey, String receivingChannel) {
        slackClient = new SlackClient(slackApiKey);
        this.receivingChannel = receivingChannel;
    }

    @Override
    public void notify(String message) {
        slackClient.chatPostMessage(receivingChannel, message);
    }
}
