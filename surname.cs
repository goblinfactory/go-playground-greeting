using System;

public static class Namer
{
    public static Func<string, string> Surname(this string lastname)
    {
        Func<string> nameit = Name (string name) =>
        {
            return name + " " + lastname;
        };
        return nameit;
    }
}

public static class Program
{
    public static void Main(string[] args)
    {
        var sn = Namer.Surname("Hemmings");
        Console.WriteLine(sn("Alan"));
    }
}
