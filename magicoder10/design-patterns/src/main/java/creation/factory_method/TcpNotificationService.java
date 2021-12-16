package creation.factory_method;

import java.net.Socket;

public class TcpNotificationService extends NotificationService {
    @Override
    public Channel makeChannel(Socket socket) {
        return new TcpChannel();
    }
}
