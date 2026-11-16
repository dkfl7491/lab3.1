#ifndef LIST_H
#define LIST_H

#include <cstddef>
#include <string>
#include <vector>

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
    void removeBefore(const std::string& target);
    void removeAfter(const std::string& target);
    void removeFront();
    void removeBack();
    void removeValue(const std::string& value);
    bool contains(const std::string& value) const;
    std::string get(std::size_t index) const;
    std::size_t size() const;
    std::vector<std::string> values() const;
};

#endif
