import React, { FC } from 'react';
import styles from './styles.module.css';

interface Props {
  children: React.ReactNodeArray;
}

const StickyHeaderLayout: FC<Props> = ({ children }) => {
  return (
    <div>
      <div className={styles.sticky}>{children[0]}</div>
      <div className={styles.body}>{children[1]}</div>
    </div>
  )
}

export default StickyHeaderLayout;