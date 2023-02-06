import React from 'react';
import PageTemplate from '../pages/PageTemplate';

export default function Welcome() {
    return (
        <PageTemplate>
            <h1 style={{textAlign: "center"}}>TenWords</h1>
            <p style={{textAlign: "center"}}>TenWords is an application that helps you learn a language with just <b>ten</b> words a day.</p>
        </PageTemplate>
    );
}