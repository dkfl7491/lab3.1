#ifndef STACK_H
#define STACK_H

#include <cstddef>
#include <string>
#include <vector>

class StackADT
{
private:
    std::string* data;
    std::size_t count;
    std::size_t capacity;

    void resize(std::size_t newCapacity);

public:
    StackADT();
    StackADT(const StackADT& other);
    StackADT& operator=(const StackADT& other);
    ~StackADT();

    void push(const std::string& value);
    std::string pop();
    std::string get() const;
    std::vector<std::string> values() const;
};

#endif
