package structure.bridge;

public class VirtualMemory implements Memory {
    private final byte[][] contents;

    VirtualMemory(int capacity) {
        contents = new byte[capacity][];
    }

    public byte[] read(int address) {
        return contents[address];
    }

    public void write(int destMemoryAddress, byte[] data) {
        contents[destMemoryAddress] = data;
    }
}
