package creation.factory_method;

import java.net.Socket;

public class PollingNotificationService extends NotificationService {
    @Override
    public Channel makeChannel(Socket socket) {
        return new PollingChannel();
    }
}
