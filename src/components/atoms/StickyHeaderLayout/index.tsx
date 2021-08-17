import React, { FC } from 'react';
import './styles.css';

interface Props {
  children: React.ReactNodeArray;
}

const StickyHeaderLayout: FC<Props> = ({ children }) => {
  return (
    <div>
      <div className='sticky'>{children[0]}</div>
      <div className='body'>{children[1]}</div>
    </div>
  )
}

export default StickyHeaderLayout;