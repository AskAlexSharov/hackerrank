//#![feature(type_alias_impl_trait)]

use serde::export::fmt::Debug;
use std::borrow::{Borrow, BorrowMut};

use std::sync::mpsc::{channel, Receiver, Sender};
use std::thread;

trait Holdable: Ord + Debug + Default + Sync + Send + Copy {}

impl<T: Ord + Debug + Default + Sync + Send + Copy> Holdable for T {}

trait Tree: std::fmt::Debug {
    fn insert(&mut self, val: i32);
    fn inorder(&self) -> Receiver<Node>;
    fn preorder(&self) -> Receiver<Node>;
    fn postorder(&self) -> Receiver<Node>;
}

impl Clone for Option<Box<Node>> {
    fn clone(&self) -> Self {
        return self.clone();
    }
}
impl Copy for Option<Box<Node>> {}

#[derive(Debug, Default, Copy, Clone)]
struct Node {
    left: Option<Box<Node>>,
    right: Option<Box<Node>>,
    data: i32,
}

impl Node {
    pub fn new(val: i32) -> Node {
        Node {
            data: val,
            ..Default::default()
        }
    }

    pub fn inorder(&self, sender: Sender<Node>) {
        sender.send(*self).unwrap()
    }
}

#[derive(Debug, Default)]
struct RecursiveTree {
    root: Option<Box<Node>>,
}

impl Tree for RecursiveTree {
    fn insert(&mut self, val: i32) {
        if self.root.is_none() {
            self.root = Some(Box::new(Node::new(val)));
            return;
        }

        let mut node = *self.root.unwrap();
        loop {
            if val == node.data {
                break;
            }

            if val < node.data {
                if node.left.is_none() {
                    let mut n = node;
                    n.left = Some(Box::new(Node::new(val)));
                    break;
                }

                node = *node.left.unwrap();
            }

            if node.right.is_none() {
                node.right = Some(Box::new(Node::new(val)));
                break;
            }

            node = *node.right.unwrap();
        }
    }

    fn inorder(&self) -> Receiver<Node> {
        let (tx, rx) = channel();
        //        thread::spawn(move || {
        //            self.root.unwrap().inorder(tx);
        //        });
        return rx;
    }
    fn preorder(&self) -> Receiver<Node> {
        panic!("not implemented yet!")
    }
    fn postorder(&self) -> Receiver<Node> {
        panic!("not implemented yet!")
    }
}

fn build_binary_search_tree(tree: &mut RecursiveTree, seed: std::vec::IntoIter<i32>) {
    seed.for_each(|x| tree.insert(x));
}

fn prepare_data_for_tree(n: i32) -> std::vec::IntoIter<i32> {
    use rand_distr::{Distribution, Normal};
    let normal = Normal::new(2.0, 3.0).unwrap();
    let mut data = vec![];
    for _i in 0..n {
        let gen = normal.sample(&mut rand::thread_rng());
        data.push(gen as i32);
    }
    data.into_iter()
}

fn main() {
    let mut tree = RecursiveTree {
        ..Default::default()
    };
    let data = prepare_data_for_tree(10);

    build_binary_search_tree(&mut tree, data);
    let rec = tree.inorder();

    dbg!(rec);
}
