#include <fstream>
#include <iostream>

int main() {
    std::ifstream file("commands.txt");
    std::string command;
    int aim = 0, depth = 0, param, pos = 0;

    while (file >> command >> param) {
        if (command == "forward") {
            pos += param;
            depth += aim*param;
        } else if (command == "up") {
            aim -= param;
        } else if (command == "down") {
            aim += param;
        }
    }

    std::cout << pos*aim << "\n";
    std::cout << pos*depth << "\n";

    return 0;
}
