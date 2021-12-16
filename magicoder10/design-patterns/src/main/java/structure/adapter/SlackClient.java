package structure.adapter;

public record SlackClient(String apiKey) {

    void chatPostMessage(String id, String message) {
        System.out.printf("Posted %s to Slack channel (%s)", message, id);
    }
}
