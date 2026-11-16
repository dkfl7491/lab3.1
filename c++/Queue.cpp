#include "Queue.h"
#include <stdexcept>

// Создание пустой очереди.
QueueADT::QueueADT()
{
    count = 0;
    capacity = 4;
    frontIndex = 0;
    data = new std::string[capacity];
}

// Создание копии очереди.
QueueADT::QueueADT(const QueueADT& other)
{
    count = other.count;
    capacity = other.capacity;
    frontIndex = 0;

    data = new std::string[capacity];

    for (std::size_t i = 0; i < count; i++)
        data[i] = other.data[(other.frontIndex + i) % other.capacity];
}

// Оператор присваивания.
QueueADT& QueueADT::operator=(const QueueADT& other)
{
    if (this == &other)
        return *this;

    std::string* newData = new std::string[other.capacity];

    for (std::size_t i = 0; i < other.count; i++)
        newData[i] = other.data[(other.frontIndex + i) % other.capacity];

    delete[] data;

    data = newData;
    count = other.count;
    capacity = other.capacity;
    frontIndex = 0;

    return *this;
}

// Освобождение памяти.
QueueADT::~QueueADT()
{
    delete[] data;
}

// Изменение вместимости очереди.
void QueueADT::resize(std::size_t newCapacity)
{
    std::string* newData = new std::string[newCapacity];

    for (std::size_t i = 0; i < count; i++)
        newData[i] = data[(frontIndex + i) % capacity];

    delete[] data;

    data = newData;
    capacity = newCapacity;
    frontIndex = 0;
}

// Добавление элемента в конец очереди.
void QueueADT::push(const std::string& value)
{
    if (count == capacity)
        resize(capacity * 2);

    std::size_t index = (frontIndex + count) % capacity;

    data[index] = value;
    count++;
}

// Удаление и возврат первого элемента очереди.
std::string QueueADT::pop()
{
    if (count == 0)
        throw std::runtime_error("Queue is empty");

    std::string value = data[frontIndex];

    frontIndex = (frontIndex + 1) % capacity;
    count--;

    if (count == 0)
        frontIndex = 0;

    return value;
}

// Возвращение первого элемента без удаления.
std::string QueueADT::get() const
{
    if (count == 0)
        throw std::runtime_error("Queue is empty");

    return data[frontIndex];
}

// Возвращение элементов очереди в порядке извлечения.
std::vector<std::string> QueueADT::values() const
{
    std::vector<std::string> result;

    for (std::size_t i = 0; i < count; i++)
        result.push_back(data[(frontIndex + i) % capacity]);

    return result;
}
