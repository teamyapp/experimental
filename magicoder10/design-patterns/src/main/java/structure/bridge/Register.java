package structure.bridge;

public interface Register {
    void write(byte[] data);

    byte[] read();
}
