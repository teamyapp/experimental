package structure.bridge;

import java.util.List;

interface Instruction {
    void execute(List<Register> registers, Memory memory, Disk disk, Output output);
}
