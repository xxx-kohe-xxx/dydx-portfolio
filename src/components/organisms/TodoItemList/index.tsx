import React , { FC } from 'react';
import styles from './styles.module.css';
import TodoItem from '../TodoItem';
import { Todo } from '../../common/Type';

interface Props {
  todos: Todo[];
}

const TodoItemList: FC<Props> = ({ todos }) => {
  return (
    <div className={styles.root}>
      { todos.map((todo, idx) => {
        return (
          <TodoItem
            key={idx}
            todo={todo}
          />
        )
      })}
    </div>
  )
}

export default TodoItemList;