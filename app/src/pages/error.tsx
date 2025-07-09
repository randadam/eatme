export function ErrorPage({ title = "Error", description = "An error occurred." }) {
    return (
        <div>
            <h2 className="text-2xl text-destructive font-bold pb-6">{title}</h2>
            <p className="text-destructive">{description}</p>
        </div>
    )
}