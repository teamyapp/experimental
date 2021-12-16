package creation.factory_method;

public class TcpChannel implements Channel {
    @Override
    public void send(String message) {
        throw new UnsupportedOperationException();
    }
}
