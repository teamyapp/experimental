package creation.factory_method;

public class LongPollChannel implements Channel {
    @Override
    public void send(String message) {
        throw new UnsupportedOperationException();
    }
}
