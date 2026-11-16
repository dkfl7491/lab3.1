#include "FullBinaryTree.h"
#include <queue>

FullBinaryTree::Node::Node(const std::string& value) : value(value) {}

// Вставляет элемент в первое свободное место при обходе дерева по уровням.
void FullBinaryTree::insert(const std::string& value)
{
    insertionOrder.push_back(value);
    auto node = std::make_unique<Node>(value);

    if (!root) { root = std::move(node); return; }

    std::queue<Node*> nodes;
    nodes.push(root.get());

    while (!nodes.empty()) {
        Node* current = nodes.front();
        nodes.pop();

        if (!current->left) { current->left = std::move(node); return; }
        nodes.push(current->left.get());

        if (!current->right) { current->right = std::move(node); return; }
        nodes.push(current->right.get());
    }
}

// Рекурсивно проверяет наличие значения в поддереве.
bool FullBinaryTree::containsNode(const Node* node, const std::string& value) const
{
    if (!node) return false;
    if (node->value == value) return true;

    return containsNode(node->left.get(), value) ||
           containsNode(node->right.get(), value);
}

// Запускает поиск значения от корневого узла.
bool FullBinaryTree::contains(const std::string& value) const
{
    return containsNode(root.get(), value);
}

// Проверяет, что каждый узел имеет либо двух потомков, либо не имеет их совсем.
bool FullBinaryTree::isFullNode(const Node* node) const
{
    if (!node) return true;
    if (!node->left && !node->right) return true;
    if (!node->left || !node->right) return false;

    return isFullNode(node->left.get()) &&
           isFullNode(node->right.get());
}

// Проверяет свойство полного бинарного дерева, начиная с корня.
bool FullBinaryTree::isFull() const
{
    return isFullNode(root.get());
}

void FullBinaryTree::collectValues(const Node* node, std::vector<std::string>& result) const
{
    if (!node) return;

    std::queue<const Node*> nodes;
    nodes.push(node);

    while (!nodes.empty()) {
        const Node* current = nodes.front();
        nodes.pop();

        result.push_back(current->value);

        if (current->left)
            nodes.push(current->left.get());

        if (current->right)
            nodes.push(current->right.get());
    }
}

std::vector<std::string> FullBinaryTree::values() const
{
    std::vector<std::string> result;
    collectValues(root.get(), result);
    return result;
}

// Возвращает элементы в порядке их добавления в дерево.
std::vector<std::string> FullBinaryTree::insertionValues() const
{
    return insertionOrder;
}
