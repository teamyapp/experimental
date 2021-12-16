package structure.decorator;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

public class BufferedTransport implements Transport {
    private final ByteArrayOutputStream byteArrayOutputStream = new ByteArrayOutputStream();

    public static void main(String[] args) {
        Transport transport = new BufferedTransport();

        byte[] encryptionKey = new byte[]{
                0x1, 0x2, 0x3, 0x4, 0x1, 0x2, 0x3, 0x4,
                0x1, 0x2, 0x3, 0x4, 0x1, 0x2, 0x3, 0x4,
                0x1, 0x2, 0x3, 0x4, 0x1, 0x2, 0x3, 0x4,
                0x1, 0x2, 0x3, 0x4, 0x1, 0x2, 0x3, 0x4,
        };
        transport = new Encryption(transport, encryptionKey);
        transport = new Compression(transport);

        try {
            transport.send(new byte[]{
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
                    0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa, 0xa,
            });
        } catch (IOException e) {
            e.printStackTrace();
        }

        byte[] data = transport.receive();

        System.out.println("Final:");
        Bytes.println(data);
    }

    @Override
    public void send(byte[] data) throws IOException {
        System.out.println("BufferedTransport(send):");
        Bytes.println(data);
        System.out.println();
        byteArrayOutputStream.write(data);
    }

    @Override
    public byte[] receive() {
        byte[] data = byteArrayOutputStream.toByteArray();
        System.out.println("BufferedTransport(receive):");
        Bytes.println(data);
        System.out.println();
        byteArrayOutputStream.reset();
        return data;
    }
}
