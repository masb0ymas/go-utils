set -e

PREV_TAG=$(git tag | sort -V | tail -n 1)
NEW_TAG="v0.0.1"

# Check if tag exists, if not create a new tag
if ! git describe --tags --abbrev=0 2>/dev/null; then
  NEW_TAG="v0.0.1"
else
  PREV_TAG=$(git describe --tags --abbrev=0)
  NEW_TAG=$(echo $PREV_TAG | awk -F. -v OFS=. '{++$NF} 1')
fi

echo -e "---------------------------------------------"
echo "| Previous Release Tag  | $PREV_TAG"
echo "| New Release Tag       | $NEW_TAG"
echo -e "---------------------------------------------\n"

# Generate changelog
echo "# Changelog" > CHANGELOG.md
echo "" >> CHANGELOG.md
echo "## $NEW_TAG" >> CHANGELOG.md
echo "" >> CHANGELOG.md

if ! git describe --tags --abbrev=0 2>/dev/null; then
  git log --pretty=format:"* %s" $HEAD >> CHANGELOG.md
else
  git log $PREV_TAG..$HEAD --pretty=format:"* %s" >> CHANGELOG.md
fi

echo "" >> CHANGELOG.md

# Commit and push
git add .
git commit -m "chore(release): bump to version $NEW_TAG"
git tag -a "$NEW_TAG" -m "Release version $NEW_TAG"
git push origin "$NEW_TAG"
git push origin main

echo -e "\nRelease a new version successfully!\n"
