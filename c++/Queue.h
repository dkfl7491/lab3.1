#ifndef QUEUE_H
#define QUEUE_H

#include <cstddef>
#include <string>
#include <vector>

class QueueADT
{
private:
    std::string* data;
    std::size_t count;
    std::size_t capacity;
    std::size_t frontIndex;

    void resize(std::size_t newCapacity);

public:
    QueueADT();
    QueueADT(const QueueADT& other);
    QueueADT& operator=(const QueueADT& other);
    ~QueueADT();

    void push(const std::string& value);
    std::string pop();
    std::string get() const;
    std::vector<std::string> values() const;
};

#endif
