import CreateButton from './create-button';


function Header({title}) {
    return <h1>{title ? title : 'Campaign Forge'}</h1>
}

export default function HomePage() {


    return (
        <div>
            <Header />
            <input/>
            <CreateButton />
        </div>
    );
}