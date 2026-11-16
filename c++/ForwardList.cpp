#include "ForwardList.h"
#include <stdexcept>
#include <utility>

SinglyList::Node::Node(const std::string& value) : value(value) {}

// Добавляет новый узел в конец списка.
void SinglyList::pushBack(const std::string& value)
{
    auto node = std::make_unique<Node>(value);
    Node* raw = node.get();
    if (!head) head = std::move(node);
    else tail->next = std::move(node);
    tail = raw;
    ++count;
}

// Добавляет новый узел в начало списка.
void SinglyList::pushFront(const std::string& value)
{
    auto node = std::make_unique<Node>(value);
    if (!head) tail = node.get();
    node->next = std::move(head);
    head = std::move(node);
    ++count;
}

// Вставляет новый узел перед узлом с указанным значением.
void SinglyList::insertBefore(const std::string& target, const std::string& value)
{
    if (!head) throw std::runtime_error("List is empty");
    if (head->value == target) { pushFront(value); return; }
    Node* previous = head.get();
    while (previous->next && previous->next->value != target)
        previous = previous->next.get();
    if (!previous->next) throw std::runtime_error("Target value not found");
    auto node = std::make_unique<Node>(value);
    node->next = std::move(previous->next);
    previous->next = std::move(node);
    ++count;
}

// Вставляет новый узел после узла с указанным значением.
void SinglyList::insertAfter(const std::string& target, const std::string& value)
{
    Node* current = head.get();
    while (current && current->value != target) current = current->next.get();
    if (!current) throw std::runtime_error("Target value not found");
    auto node = std::make_unique<Node>(value);
    node->next = std::move(current->next);
    current->next = std::move(node);
    if (!current->next->next) tail = current->next.get();
    ++count;
}

// Удаляет первый узел списка.
void SinglyList::removeFront()
{
    if (!head) throw std::runtime_error("List is empty");
    head = std::move(head->next);
    --count;
    if (!head) tail = nullptr;
}

// Удаляет последний узел списка.
void SinglyList::removeBack()
{
    if (!head) throw std::runtime_error("List is empty");
    if (head.get() == tail) { head.reset(); tail = nullptr; count = 0; return; }
    Node* current = head.get();
    while (current->next.get() != tail) current = current->next.get();
    current->next.reset();
    tail = current;
    --count;
}

// Находит и удаляет узел с указанным значением.
void SinglyList::removeValue(const std::string& value)
{
    if (!head) throw std::runtime_error("List is empty");
    if (head->value == value) { removeFront(); return; }
    Node* previous = head.get();
    while (previous->next && previous->next->value != value)
        previous = previous->next.get();
    if (!previous->next) throw std::runtime_error("Value not found");
    if (previous->next.get() == tail) tail = previous;
    previous->next = std::move(previous->next->next);
    --count;
}

// Удаляет узел перед узлом с указанным значением.
void SinglyList::removeBefore(const std::string& target)
{
    if (!head) throw std::runtime_error("List is empty");
    if (head->value == target) throw std::runtime_error("No element before target");
    if (head->next && head->next->value == target) { removeFront(); return; }
    Node* prevPrev = head.get();
    while (prevPrev->next && prevPrev->next->next && prevPrev->next->next->value != target)
        prevPrev = prevPrev->next.get();
    if (!prevPrev->next || !prevPrev->next->next) throw std::runtime_error("Target value not found");
    prevPrev->next = std::move(prevPrev->next->next);
    --count;
}

// Удаляет узел после узлом с указанным значением.
void SinglyList::removeAfter(const std::string& target)
{
    if (!head) throw std::runtime_error("List is empty");
    Node* current = head.get();
    while (current && current->value != target) current = current->next.get();
    if (!current) throw std::runtime_error("Target value not found");
    if (!current->next) throw std::runtime_error("No element after target");
    if (current->next.get() == tail) tail = current;
    current->next = std::move(current->next->next);
    --count;
}

// Проверяет наличие значения в списке.
bool SinglyList::contains(const std::string& value) const
{
    for (Node* current = head.get(); current; current = current->next.get())
        if (current->value == value) return true;
    return false;
}

// Возвращает значение узла по индексу.
std::string SinglyList::get(std::size_t index) const
{
    if (index >= count) throw std::out_of_range("Invalid list index");
    Node* current = head.get();
    for (std::size_t i = 0; i < index; ++i) current = current->next.get();
    return current->value;
}

std::size_t SinglyList::size() const { return count; }

// Возвращает значения всех узлов в порядке следования.
std::vector<std::string> SinglyList::values() const
{
    std::vector<std::string> result;
    for (Node* current = head.get(); current; current = current->next.get())
        result.push_back(current->value);
    return result;
}

