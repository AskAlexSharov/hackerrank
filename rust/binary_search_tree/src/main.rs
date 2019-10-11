//#![feature(type_alias_impl_trait)]

use serde::export::fmt::Debug;
use std::borrow::{Borrow, BorrowMut};

type Visitor<T> = Box<Fn(&Node<T>) -> bool>;

trait Holdable: Ord + Debug + Default {}

impl<T: Ord + Debug + Default> Holdable for T {}

trait Tree<T: Holdable>: std::fmt::Debug {
    fn insert(&mut self, val: T);
    fn inorder(&mut self, visitor: Visitor<T>);
    fn preorder(&mut self, visitor: Visitor<T>);
    fn postorder(&mut self, visitor: Visitor<T>);
}

type NodeId = usize;

#[derive(Debug, Default)]
struct Node<T: Holdable> {
    left: Option<NodeId>,
    right: Option<NodeId>,
    data: T,
}

impl<T: Holdable> From<T> for Node<T> {
    fn from(val: T) -> Self {
        Node::new(val)
    }
}

impl<T: Holdable> Node<T> {
    pub fn new(val: T) -> Node<T> {
        Node {
            data: val,
            ..Default::default()
        }
    }
}

#[derive(Debug, Default)]
struct Arena<T: Holdable> {
    pub(crate) nodes: Vec<Node<T>>,
}

impl<T: Holdable> Arena<T> {
    fn add(&mut self, val: T) -> NodeId {
        let k = self.nodes.len();
        self.nodes.insert(k, Node::new(val));
        k
    }

    pub fn get(&self, id: NodeId) -> &Node<T> {
        self.nodes.get(id).unwrap()
    }

    pub fn get_mut(&mut self, id: NodeId) -> &mut Node<T> {
        self.nodes.get_mut(id).unwrap()
    }
}

#[derive(Debug, Default)]
struct MorrisTree<T: Holdable> {
    root: Option<NodeId>,
    arena: Arena<T>,
}

fn left_right_most_child<T: Holdable>(node_id: NodeId, arena: &mut Arena<T>) -> Option<NodeId> {
    let mut child_id = arena.get(node_id).left?;
    let child = arena.get(child_id);

    while child.right.is_some() && child.right.unwrap() != node_id {
        child_id = child.right.unwrap();
    }

    Some(child_id)
}

impl<T: Holdable> Tree<T> for MorrisTree<T> {
    fn insert(&mut self, val: T) {
        let arena = self.arena.borrow_mut();

        if self.root.is_none() {
            self.root = Some(arena.add(val));
            return;
        }

        let mut node_id = self.root.unwrap();
        loop {
            let node = arena.get(node_id);

            if val == node.data {
                break;
            }

            if val < node.data {
                if node.left.is_none() {
                    arena.get_mut(node_id).left = Some(arena.add(val));
                    break;
                }

                node_id = node.left.unwrap();
            }

            if node.right.is_none() {
                arena.get_mut(node_id).right = Some(arena.add(val));
                break;
            }

            node_id = node.right.unwrap();
        }
    }

    fn inorder(&mut self, visitor: Visitor<T>) {
        let mut node_id = self.root;
        let arena = self.arena.borrow_mut();

        while node_id.is_some() {
            let node = arena.get_mut(node_id.unwrap());
            if node.left.is_none() {
                if !visitor(node) {
                    break;
                }

                node_id = node.right;
                continue;
            }

            // left_right_most_child
            let mut pre_id = node.left.unwrap();
            let pre = arena.get_mut(pre_id);
            while pre.right.is_some() && pre.right.unwrap() != pre_id {
                pre_id = pre.right.unwrap();
            }

            let pre_node = arena.get_mut(pre_id);

            if pre_node.right.is_none() || pre_node.right.unwrap() != node_id.unwrap() {
                pre_node.right = Some(node_id.unwrap());
                node_id = node.left;
                continue;
            }

            pre_node.right.take();

            if !visitor(node) {
                break;
            }

            node_id = node.right;
        }
        panic!("not implemented yet!")
    }
    fn preorder(&mut self, visitor: Visitor<T>) {
        panic!("not implemented yet!")
    }
    fn postorder(&mut self, visitor: Visitor<T>) {
        panic!("not implemented yet!")
    }
}

fn build_binary_search_tree<T: Holdable>(tree: &mut MorrisTree<T>, seed: std::vec::IntoIter<T>) {
    seed.for_each(|x| tree.insert(x));
}

fn prepare_data_for_tree(n: i32) -> std::vec::IntoIter<i32> {
    use rand_distr::{Distribution, Normal};
    let normal = Normal::new(2.0, 3.0).unwrap();
    let mut data = vec![];
    for i in 0..n {
        let gen = normal.sample(&mut rand::thread_rng());
        data.push(gen as i32);
    }
    data.into_iter()
}

fn main() {
    let mut tree = MorrisTree::<i32> {
        arena: Arena { nodes: vec![] },
        ..Default::default()
    };
    let data = prepare_data_for_tree(10);

    build_binary_search_tree(&mut tree, data);
    tree.inorder(Box::new(|node| {
        dbg!(node);
        true
    }));

    //    dbg!(tree);
}
