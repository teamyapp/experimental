package structure.bridge;

import java.nio.ByteBuffer;
import java.util.List;

public record ShowInstruction(int memoryAddress) implements Instruction {
    @Override
    public void execute(List<Register> registers, Memory memory, Disk disk, Output output) {
        output.write(ByteBuffer.wrap(memory.read(memoryAddress)).getInt());
    }
}
