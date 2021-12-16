package creation.factory_method;

import java.net.Socket;
import java.util.HashMap;
import java.util.Map;

public abstract class NotificationService {
    private final Map<String, Channel> channels = new HashMap<>();

    public void onUserConnect(String userId, Socket client) {
        channels.put(userId, makeChannel(client));
    }

    public abstract Channel makeChannel(Socket socket);

    public void notifyUser(String toUserId, String message) {
        if (!channels.containsKey(toUserId)) {
            return;
        }

        Channel channel = channels.get(toUserId);
        channel.send(message);
    }
}
