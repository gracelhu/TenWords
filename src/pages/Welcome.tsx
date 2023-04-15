import React from 'react';
import PageTemplate from '../pages/PageTemplate';

export default function Welcome() {
    return (
        <PageTemplate>
            <h1 style={{textAlign: "center"}}>TenWords</h1>
            <p style={{textAlign: "center"}}>TenWords is an application that helps you to learn a language with just <b>ten</b> words a day.</p>
            <p style={{textAlign: "center"}}>Head over to the 'Your Words' tab to start learning!</p>
        </PageTemplate>
    );
}