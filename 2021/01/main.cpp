#include <fstream>
#include <iostream>

#include <boost/circular_buffer.hpp>

int countIncreases(int windowSize) {
    std::ifstream file("measurements.txt");
    boost::circular_buffer<int> cb(windowSize);
    int measurement, result = 0;

    while (file >> measurement) {
        cb.push_front(measurement);
        if (cb.full() && cb.front() > cb.back()) {
            result++;
        }
    }
    return result;
}

int main() {
    std::cout << countIncreases(2) << "\n";
    std::cout << countIncreases(4) << "\n";

    return 0;
}
