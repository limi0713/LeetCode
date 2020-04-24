#include <iostream>
using namespace std;

bool myfunc(string s, int i, string p, int j)
{
    if (i < s.length() && j >= p.length())
    {
        return false;
    }
    if (i == s.length() && j == p.length())
    {
        return true;
    }
    if (i == s.length())
    {
        if ((p.length() - j) % 2 == 1)
        {
            return false;
        }
        while (j + 1 < p.length())
        {
            if (p[j + 1] == '*')
            {
                j += 2;
            }
            else
            {
                return false;
            }
        }
        return true;
    }
    if ((j + 1) < p.length() && p[j + 1] == '*')
    {
        if (s[i] == p[j] || p[j] == '.')
        {
            return myfunc(s, i + 1, p, j) || myfunc(s, i, p, j + 2);
        }
        else
        {
            return myfunc(s, i, p, j + 2);
        }
    }
    else if (s[i] == p[j] || p[j] == '.')
    {
        return myfunc(s, i + 1, p, j + 1);
    }
    else if (s[i] != p[j])
    {
        return false;
    }
    return false;
}

bool isMatch(string s, string p)
{
    return myfunc(s, 0, p, 0);
}

int main()
{
    cout << endl
         << isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c") << endl;
}