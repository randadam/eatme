export function NotFoundPage({ title = "Not Found", description = "The requested resource does not exist." }) {
    return (
        <div className="text-center py-10 text-muted-foreground">
            <h2 className="text-xl font-semibold mb-2">{title}</h2>
            <p>{description}</p>
        </div>
    );
}
    