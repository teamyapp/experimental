package structure.bridge;

import java.nio.ByteBuffer;
import java.util.Arrays;
import java.util.List;

public record Machine(List<Register> registers, Memory memory, Disk disk, Output output) {
    public static void main(String[] args) {
        List<Register> registers = Arrays.asList(
                new VirtualRegister(),
                new VirtualRegister(),
                new VirtualRegister()
        );

        VirtualMemory virtualMemory = new VirtualMemory(100);
        VirtualDisk virtualDisk = new VirtualDisk();
        Terminal output = new Terminal();

        Machine machine = new Machine(registers, virtualMemory, virtualDisk, output);

        virtualMemory.write(0, ByteBuffer.allocate(4).putInt(12).array());
        virtualMemory.write(1, ByteBuffer.allocate(4).putInt(58).array());

        machine.execute(Arrays.asList(
                new LoadInstruction(0, 0),
                new LoadInstruction(1, 1),
                new AddInstruction(0, 1, 2),
                new SaveInstruction(2, 3),
                new ShowInstruction(3)
        ));
    }

    public void execute(List<Instruction> instructions) {
        for (Instruction instruction : instructions) {
            instruction.execute(registers, memory, disk, output);
        }
    }
}
