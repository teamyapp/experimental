package structure.decorator;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.zip.GZIPInputStream;
import java.util.zip.GZIPOutputStream;

public class Compression extends TransportDecorator {
    Compression(Transport transport) {
        super(transport);
    }

    @Override
    public void send(byte[] data) throws IOException {
        System.out.println("Compression(send/input):");
        Bytes.println(data);
        System.out.println();

        ByteArrayOutputStream byteArrayOutputStream = null;
        GZIPOutputStream gzipOutputStream = null;

        try {
            byteArrayOutputStream = new ByteArrayOutputStream();
            gzipOutputStream = new GZIPOutputStream(byteArrayOutputStream);
            gzipOutputStream.write(data);
        } finally {
            gzipOutputStream.close();
            byteArrayOutputStream.close();
        }

        byte[] compressedData = byteArrayOutputStream.toByteArray();
        System.out.println("Compression(send/output):");
        Bytes.println(compressedData);
        System.out.println();

        transport.send(compressedData);
    }

    @Override
    public byte[] receive() {
        byte[] compressedData = transport.receive();
        System.out.println("Compression(receive/input):");
        Bytes.println(compressedData);
        System.out.println();

        try {
            ByteArrayInputStream byteArrayInputStream = new ByteArrayInputStream(compressedData);
            GZIPInputStream gzipInputStream = new GZIPInputStream(byteArrayInputStream);
            byte[] data = gzipInputStream.readAllBytes();

            System.out.println("Compression(receive/output):");
            Bytes.println(data);
            System.out.println();

            return data;
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
