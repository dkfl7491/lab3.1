#ifndef FULL_BINARY_TREE_H
#define FULL_BINARY_TREE_H

#include <memory>
#include <string>
#include <vector>

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

    void collectValues(
        const Node* node,
        std::vector<std::string>& result
    ) const;

public:
    void insert(const std::string& value);
    bool contains(const std::string& value) const;
    bool isFull() const;
    std::vector<std::string> values() const;
    std::vector<std::string> insertionValues() const;
};

#endif
