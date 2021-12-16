package creation.factory_method;

public class PollingChannel implements Channel {
    @Override
    public void send(String message) {
        throw new UnsupportedOperationException();
    }
}
