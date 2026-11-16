#include "List.h"
#include <stdexcept>

DoublyList::Node::Node(const std::string& value) : value(value) {}

// Освобождает память всех узлов списка.
DoublyList::~DoublyList()
{
    Node* current = head;
    while (current) { Node* next = current->next; delete current; current = next; }
}

// Ищет узел с заданным значением.
DoublyList::Node* DoublyList::findNode(const std::string& value) const
{
    Node* current = head;
    while (current) {
        if (current->value == value) return current;
        current = current->next;
    }
    return nullptr;
}

// Добавляет новый узел в конец списка.
void DoublyList::pushBack(const std::string& value)
{
    Node* node = new Node(value);
    node->prev = tail;
    if (tail) tail->next = node;
    else head = node;
    tail = node;
    ++count;
}

// Добавляет новый узел в начало списка.
void DoublyList::pushFront(const std::string& value)
{
    Node* node = new Node(value);
    node->next = head;
    if (head) head->prev = node;
    else tail = node;
    head = node;
    ++count;
}

// Вставляет новый узел перед первым найденным узлом.
void DoublyList::insertBefore(const std::string& target, const std::string& value)
{
    Node* targetNode = findNode(target);
    if (!targetNode) throw std::runtime_error("Target value not found");
    if (targetNode == head) { pushFront(value); return; }
    Node* node = new Node(value);
    node->prev = targetNode->prev;
    node->next = targetNode;
    targetNode->prev->next = node;
    targetNode->prev = node;
    ++count;
}

// Вставляет новый узел после первого найденного узла.
void DoublyList::insertAfter(const std::string& target, const std::string& value)
{
    Node* targetNode = findNode(target);
    if (!targetNode) throw std::runtime_error("Target value not found");
    if (targetNode == tail) { pushBack(value); return; }
    Node* node = new Node(value);
    node->prev = targetNode;
    node->next = targetNode->next;
    targetNode->next->prev = node;
    targetNode->next = node;
    ++count;
}

// Удаляет узел и связывает между собой его соседей.
void DoublyList::removeNode(Node* node)
{
    if (!node) throw std::runtime_error("Value not found");
    if (node->prev) node->prev->next = node->next;
    else head = node->next;
    if (node->next) node->next->prev = node->prev;
    else tail = node->prev;
    delete node;
    --count;
}

void DoublyList::removeFront()
{
    if (!head) throw std::runtime_error("List is empty");
    removeNode(head);
}

void DoublyList::removeBack()
{
    if (!tail) throw std::runtime_error("List is empty");
    removeNode(tail);
}

void DoublyList::removeValue(const std::string& value) { removeNode(findNode(value)); }

// Удаляет узел перед узлом с указанным значением.
void DoublyList::removeBefore(const std::string& target)
{
    if (!head) throw std::runtime_error("List is empty");
    Node* targetNode = findNode(target);
    if (!targetNode) throw std::runtime_error("Target value not found");
    if (!targetNode->prev) throw std::runtime_error("No element before target");
    removeNode(targetNode->prev);
}

// Удаляет узел после узла с указанным значением.
void DoublyList::removeAfter(const std::string& target)
{
    if (!head) throw std::runtime_error("List is empty");
    Node* targetNode = findNode(target);
    if (!targetNode) throw std::runtime_error("Target value not found");
    if (!targetNode->next) throw std::runtime_error("No element after target");
    removeNode(targetNode->next);
}

bool DoublyList::contains(const std::string& value) const { return findNode(value) != nullptr; }

// Возвращает значение узла по его порядковому индексу.
std::string DoublyList::get(std::size_t index) const
{
    if (index >= count) throw std::out_of_range("Invalid list index");
    Node* current = head;
    for (std::size_t i = 0; i < index; ++i) current = current->next;
    return current->value;
}

std::size_t DoublyList::size() const { return count; }

// Формирует вектор со значениями всех узлов в порядке обхода.
std::vector<std::string> DoublyList::values() const
{
    std::vector<std::string> result;
    for (Node* current = head; current; current = current->next)
        result.push_back(current->value);
    return result;
}

