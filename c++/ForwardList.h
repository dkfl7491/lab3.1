#ifndef FORWARD_LIST_H
#define FORWARD_LIST_H

#include <cstddef>
#include <memory>
#include <string>
#include <vector>

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
