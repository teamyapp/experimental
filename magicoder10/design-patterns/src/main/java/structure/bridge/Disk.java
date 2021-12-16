package structure.bridge;

public interface Disk {
    void save(String filePath, byte[] data);

    byte[] read(String filePath);
}
