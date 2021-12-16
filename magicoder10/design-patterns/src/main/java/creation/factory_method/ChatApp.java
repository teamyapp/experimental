package creation.factory_method;

public class ChatApp {
    private final NotificationService notificationService;

    public ChatApp(NotificationService notificationService) {
        this.notificationService = notificationService;
    }

    public static void main(String[] args) {
        ChatApp tcpChatApp = new ChatApp(new TcpNotificationService());
        tcpChatApp.sendMessage("userA", "Factory method");

        ChatApp longPollChatApp = new ChatApp(new LongPollNotificationService());
        longPollChatApp.sendMessage("userA", "Factory method");

        ChatApp pollingChatApp = new ChatApp(new PollingNotificationService());
        pollingChatApp.sendMessage("userA", "Factory method");
    }

    public void sendMessage(String toUserId, String message) {
        notificationService.notifyUser(toUserId, message);
    }
}
