package structure.bridge;

import java.nio.ByteBuffer;
import java.util.List;

public record AddInstruction(
        int num1registerIndex,
        int num2registerIndex,
        int destRegisterIndex) implements Instruction {
    @Override
    public void execute(List<Register> registers, Memory memory, Disk disk, Output output) {
        byte[] bytes1 = registers.get(num1registerIndex).read();
        byte[] bytes2 = registers.get(num2registerIndex).read();
        int num1 = ByteBuffer.wrap(bytes1).getInt();
        int num2 = ByteBuffer.wrap(bytes2).getInt();
        registers.get(destRegisterIndex).write(ByteBuffer.allocate(4).putInt(num1 + num2).array());
    }
}
