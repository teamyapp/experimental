package structure.bridge;

public class VirtualRegister implements Register {
    private byte[] data;

    @Override
    public void write(byte[] data) {
        this.data = data;
    }

    @Override
    public byte[] read() {
        return this.data;
    }
}
