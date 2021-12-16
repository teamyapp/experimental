package creation.factory_method;

import java.net.Socket;

public class LongPollNotificationService extends NotificationService {
    @Override
    public Channel makeChannel(Socket socket) {
        return new LongPollChannel();
    }
}
