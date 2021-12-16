package structure.bridge;

import java.util.List;

public record SaveInstruction(int srcRegisterIndex, int destMemoryAddress) implements Instruction {
    @Override
    public void execute(List<Register> registers, Memory memory, Disk disk, Output output) {
        memory.write(destMemoryAddress, registers.get(srcRegisterIndex).read());
    }
}
