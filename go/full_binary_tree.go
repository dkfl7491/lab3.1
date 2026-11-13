package main

import "container/list"

// Узел бинарного дерева.
type TreeNode struct {
    Value string
    Left  *TreeNode
    Right *TreeNode
}

// Бинарное дерево и порядок добавления элементов.
type FullBinaryTree struct {
    Root           *TreeNode
    InsertionOrder []string
}

// Добавление узла на первое свободное место по уровням.
func (t *FullBinaryTree) Insert(value string) {
    t.InsertionOrder = append(t.InsertionOrder, value)
    node := &TreeNode{Value: value}

    if t.Root == nil {
        t.Root = node
        return
    }

    // Очередь узлов для обхода по уровням.
    nodes := list.New()
    nodes.PushBack(t.Root)

    for nodes.Len() > 0 {
        element := nodes.Front()
        current := element.Value.(*TreeNode)
        nodes.Remove(element)

        if current.Left == nil {
            current.Left = node
            return
        }

        nodes.PushBack(current.Left)

        if current.Right == nil {
            current.Right = node
            return
        }

        nodes.PushBack(current.Right)
    }
}

// Поиск значения обходом по уровням.
func (t *FullBinaryTree) Contains(value string) bool {
    if t.Root == nil {
        return false
    }

    nodes := list.New()
    nodes.PushBack(t.Root)

    for nodes.Len() > 0 {
        element := nodes.Front()
        current := element.Value.(*TreeNode)
        nodes.Remove(element)

        if current.Value == value {
            return true
        }

        if current.Left != nil {
            nodes.PushBack(current.Left)
        }

        if current.Right != nil {
            nodes.PushBack(current.Right)
        }
    }

    return false
}

// Проверка узла и его поддеревьев на полноту.
func isFullNode(node *TreeNode) bool {
    if node == nil {
        return true
    }

    // Узел без потомков.
    if node.Left == nil && node.Right == nil {
        return true
    }

    // Узел только с одним потомком.
    if node.Left == nil || node.Right == nil {
        return false
    }

    // Рекурсивная проверка обоих поддеревьев.
    return isFullNode(node.Left) && isFullNode(node.Right)
}

// Проверка полноты всего дерева.
func (t *FullBinaryTree) IsFull() bool {
    return isFullNode(t.Root)
}

// Получение значений в прямом порядке обхода.
func (t *FullBinaryTree) Values() []string {
    result := make([]string, 0)

    // Рекурсивный обход: корень, левое и правое поддеревья.
    var walk func(*TreeNode)

    walk = func(node *TreeNode) {
        if node == nil {
            return
        }

        result = append(result, node.Value)
        walk(node.Left)
        walk(node.Right)
    }

    walk(t.Root)
    return result
}	
