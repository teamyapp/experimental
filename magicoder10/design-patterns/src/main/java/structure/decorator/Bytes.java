package structure.decorator;

public class Bytes {
    static void println(byte[] data) {
        for (byte oneByte : data) {
            System.out.printf("0x%02X ", oneByte);
        }

        System.out.println();
    }
}
