package structure.decorator;

import java.io.IOException;

public interface Transport {
    void send(byte[] data) throws IOException;

    byte[] receive();
}
