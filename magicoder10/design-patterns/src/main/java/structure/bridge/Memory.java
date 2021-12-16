package structure.bridge;

public interface Memory {
    byte[] read(int address);

    void write(int destMemoryAddress, byte[] data);
}
