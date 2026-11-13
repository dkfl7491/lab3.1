#ifndef STRUCTURES_H
#define STRUCTURES_H

#include <cstddef>
#include <memory>
#include <queue>
#include <string>
#include <vector>

class ArrayADT
{
private:
    std::vector<std::string> data;

public:
    void push(const std::string& value);
    void insert(std::size_t index, const std::string& value);
    std::string get(std::size_t index) const;
    void remove(std::size_t index);
    void set(std::size_t index, const std::string& value);
    std::size_t size() const;
    std::vector<std::string> values() const;
};

class SinglyList
{
private:
    struct Node
    {
        std::string value;
        std::unique_ptr<Node> next;

        explicit Node(const std::string& value);
    };

    std::unique_ptr<Node> head;
    Node* tail = nullptr;
    std::size_t count = 0;

public:
    void pushBack(const std::string& value);
    void pushFront(const std::string& value);
    void insertBefore(const std::string& target, const std::string& value);
    void insertAfter(const std::string& target, const std::string& value);
    void removeFront();
    void removeBack();
    void removeValue(const std::string& value);
    bool contains(const std::string& value) const;
    std::string get(std::size_t index) const;
    std::size_t size() const;
    std::vector<std::string> values() const;
};

class DoublyList
{
private:
    struct Node
    {
        std::string value;
        Node* prev = nullptr;
        Node* next = nullptr;

        explicit Node(const std::string& value);
    };

    Node* head = nullptr;
    Node* tail = nullptr;
    std::size_t count = 0;

    Node* findNode(const std::string& value) const;
    void removeNode(Node* node);

public:
    DoublyList() = default;
    ~DoublyList();

    DoublyList(const DoublyList&) = delete;
    DoublyList& operator=(const DoublyList&) = delete;

    void pushBack(const std::string& value);
    void pushFront(const std::string& value);
    void insertBefore(const std::string& target, const std::string& value);
    void insertAfter(const std::string& target, const std::string& value);
    void removeFront();
    void removeBack();
    void removeValue(const std::string& value);
    bool contains(const std::string& value) const;
    std::string get(std::size_t index) const;
    std::size_t size() const;
    std::vector<std::string> values() const;
};

class StackADT
{
private:
    std::vector<std::string> data;

public:
    void push(const std::string& value);
    std::string pop();
    std::string get() const;
    std::vector<std::string> values() const;
};

class QueueADT
{
private:
    std::queue<std::string> data;

public:
    void push(const std::string& value);
    std::string pop();
    std::string get() const;
    std::vector<std::string> values() const;
};

class FullBinaryTree
{
private:
    struct Node
    {
        std::string value;
        std::unique_ptr<Node> left;
        std::unique_ptr<Node> right;

        explicit Node(const std::string& value);
    };

    std::unique_ptr<Node> root;
    std::vector<std::string> insertionOrder;

    bool containsNode(const Node* node, const std::string& value) const;
    bool isFullNode(const Node* node) const;
    void collectValues(const Node* node,
                       std::vector<std::string>& result) const;

public:
    void insert(const std::string& value);
    bool contains(const std::string& value) const;
    bool isFull() const;
    std::vector<std::string> values() const;
    std::vector<std::string> insertionValues() const;
};

#endif
