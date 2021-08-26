import React , { FC, MouseEventHandler } from 'react';
import styles from './styles.module.css';
import StickyHeaderLayout from '../../atoms/StickyHeaderLayout';
import Card, { CardHeader } from '../../atoms/Card';
import InputForm from '../../molecules/InputForm';
import TodoItemList from '../../organisms/TodoItemList';
import { Todo } from '../../common/Type';

interface Props {
  todos: Todo[];
  onSubmit: MouseEventHandler<HTMLImageElement>
}

const TodoAppTemplate: FC<Props>= ({
  todos,
  onSubmit,
}) => {
  return (
      <StickyHeaderLayout>
        <header className={styles.header}>
          <Card>
            <CardHeader>ToDo App</CardHeader>
            <InputForm onSubmit={onSubmit}/>
          </Card>
        </header>
        <main className={styles.main}>
          <TodoItemList todos={todos}/>
        </main>
      </StickyHeaderLayout>
  )
}

export default TodoAppTemplate;