package structure.decorator;

public abstract class TransportDecorator implements Transport {
    protected final Transport transport;

    TransportDecorator(Transport transport) {
        this.transport = transport;
    }
}
