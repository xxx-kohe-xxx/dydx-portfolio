import React from 'react';
import { Todo } from '../common/Type';
import TodoAppTemplate from '../templates/TodoAppTemplate';

// TODO:バックエンドからデータを取得する
const todos: Todo[] = [
  {
    id: 1,
    title: 'task1',
    done: false,
  },
  {
    id: 2,
    title: 'task2',
    done: false,
  },
  {
    id: 3,
    title: 'task3',
    done: false,
  },
];

const onSubmit = () => {
  // TODO:タスク登録処理
}

const TodoAppPage = () => {
  return (
    <TodoAppTemplate
      todos={todos}
      onSubmit={onSubmit}
    />
  )
}

export default TodoAppPage;