package structure.bridge;

import java.util.List;

public record LoadInstruction(int memoryAddress, int destRegisterIndex) implements Instruction {
    @Override
    public void execute(List<Register> registers, Memory memory, Disk disk, Output output) {
        registers.get(destRegisterIndex).write(memory.read(memoryAddress));
    }
}
