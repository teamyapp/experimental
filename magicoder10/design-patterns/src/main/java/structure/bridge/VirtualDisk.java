package structure.bridge;

import java.util.HashMap;
import java.util.Map;

public class VirtualDisk implements Disk {
    private final Map<String, byte[]> storage = new HashMap<>();

    public void save(String filePath, byte[] data) {
        storage.put(filePath, data);
    }

    public byte[] read(String filePath) {
        return storage.get(filePath);
    }
}
