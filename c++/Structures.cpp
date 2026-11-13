    {
        result.push_back(copy.front());
        copy.pop();
    }

    return result;
}

// Full Binary Tree

// Создаёт узел бинарного дерева.
FullBinaryTree::Node::Node(const std::string& value) : value(value) {}

// Вставляет узел на первое свободное место при обходе по уровням.
void FullBinaryTree::insert(const std::string& value)
{
    insertionOrder.push_back(value);

    auto node = std::make_unique<Node>(value);

    if (!root)
    {
        root = std::move(node);
        return;
    }

    std::queue<Node*> nodes;
    nodes.push(root.get());

    while (!nodes.empty())
    {
        Node* current = nodes.front();
        nodes.pop();

        if (!current->left)
        {
            current->left = std::move(node);
            return;
        }

        nodes.push(current->left.get());

        if (!current->right)
        {
            current->right = std::move(node);
            return;
        }

        nodes.push(current->right.get());
    }
}

// Рекурсивно ищет значение в поддереве.
bool FullBinaryTree::containsNode(const Node* node,
                                  const std::string& value) const
{
    if (!node)
        return false;

    if (node->value == value)
        return true;

    return containsNode(node->left.get(), value) ||
           containsNode(node->right.get(), value);
}

// Проверяет наличие значения в дереве.
bool FullBinaryTree::contains(const std::string& value) const
{
    return containsNode(root.get(), value);
}

// Проверяет, что у каждого узла ноль или два потомка.
bool FullBinaryTree::isFullNode(const Node* node) const
{
    if (!node)
        return true;

    if (!node->left && !node->right)
        return true;

    if (!node->left || !node->right)
        return false;

    return isFullNode(node->left.get()) &&
           isFullNode(node->right.get());
}

// Проверяет полноту всего дерева.
bool FullBinaryTree::isFull() const
{
    return isFullNode(root.get());
}

// Собирает значения дерева в прямом обходе: корень, левое, правое.
void FullBinaryTree::collectValues(
    const Node* node,
    std::vector<std::string>& result) const
{
    if (!node)
        return;

    result.push_back(node->value);
    collectValues(node->left.get(), result);
    collectValues(node->right.get(), result);
}

// Возвращает значения дерева в прямом обходе.
std::vector<std::string> FullBinaryTree::values() const
{
    std::vector<std::string> result;
    collectValues(root.get(), result);
    return result;
}

// Возвращает значения в порядке их добавления для сохранения дерева.
std::vector<std::string> FullBinaryTree::insertionValues() const
{
    return insertionOrder;
}
