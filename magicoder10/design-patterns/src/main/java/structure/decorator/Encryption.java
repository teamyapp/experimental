package structure.decorator;

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import javax.crypto.spec.SecretKeySpec;
import java.io.IOException;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;

public class Encryption extends TransportDecorator {
    private final SecretKeySpec key;

    Encryption(Transport transport, byte[] key) {
        super(transport);
        this.key = new SecretKeySpec(key, "AES");
    }

    @Override
    public void send(byte[] data) throws IOException {
        System.out.println("Encryption(send/input):");
        Bytes.println(data);
        System.out.println();

        Cipher cipher = GetCipher(Cipher.ENCRYPT_MODE);

        byte[] encryptedData;
        try {
            encryptedData = cipher.doFinal(data);
        } catch (IllegalBlockSizeException | BadPaddingException e) {
            throw new RuntimeException(e.toString());
        }

        System.out.println("Encryption(send/output):");
        Bytes.println(encryptedData);
        System.out.println();
        transport.send(encryptedData);
    }

    @Override
    public byte[] receive() {
        Cipher cipher = GetCipher(Cipher.DECRYPT_MODE);

        byte[] encryptedData = transport.receive();
        System.out.println("Encryption(receive/input):");
        Bytes.println(encryptedData);
        System.out.println();

        try {
            byte[] data = cipher.doFinal(encryptedData);
            System.out.println("Encryption(receive):");
            Bytes.println(data);
            System.out.println();
            return data;
        } catch (IllegalBlockSizeException | BadPaddingException e) {
            throw new RuntimeException(e.toString());
        }
    }

    private Cipher GetCipher(int opMode) {
        try {
            Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5PADDING");
            cipher.init(opMode, key);
            return cipher;
        } catch (NoSuchPaddingException | NoSuchAlgorithmException | InvalidKeyException e) {
            throw new RuntimeException(e.toString());
        }
    }
}