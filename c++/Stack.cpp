#include "Stack.h"
#include <stdexcept>

// Конструктор
StackADT::StackADT()
{
    capacity = 10;
    count = 0;
    data = new std::string[capacity];
}

// Конструктор копирования
StackADT::StackADT(const StackADT& other)
{
    capacity = other.capacity;
    count = other.count;

    data = new std::string[capacity];

    for (std::size_t i = 0; i < count; i++)
        data[i] = other.data[i];
}

// Оператор присваивания
StackADT& StackADT::operator=(const StackADT& other)
{
    if (this != &other)
    {
        delete[] data;

        capacity = other.capacity;
        count = other.count;

        data = new std::string[capacity];

        for (std::size_t i = 0; i < count; i++)
            data[i] = other.data[i];
    }

    return *this;
}

// Деструктор
StackADT::~StackADT()
{
    delete[] data;
}

// Изменение размера динамического массива
void StackADT::resize(std::size_t newCapacity)
{
    std::string* newData = new std::string[newCapacity];

    for (std::size_t i = 0; i < count; i++)
        newData[i] = data[i];

    delete[] data;

    data = newData;
    capacity = newCapacity;
}

// Добавление элемента на вершину стека
void StackADT::push(const std::string& value)
{
    if (count == capacity)
        resize(capacity * 2);

    data[count] = value;
    count++;
}

// Удаление верхнего элемента стека и возврат его значения
std::string StackADT::pop()
{
    if (count == 0)
        throw std::runtime_error("Stack is empty");

    count--;

    return data[count];
}

// Получение верхнего элемента без его удаления
std::string StackADT::get() const
{
    if (count == 0)
        throw std::runtime_error("Stack is empty");

    return data[count - 1];
}

// Получение всех элементов стека
std::vector<std::string> StackADT::values() const
{
    std::vector<std::string> result;

    for (std::size_t i = count; i > 0; i--)
        result.push_back(data[i - 1]);

    return result;
}
